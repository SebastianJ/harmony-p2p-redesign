## Overview
Harmony's P2P layer is the central component for faciliating communication between nodes in Harmony's architecture.

*With the goal of <8s finality, increased limit of external validator slots, fast syncing and resharding, now might be a suitable time to re-evaluate and potentially redesign the current P2P layer.*

## Current issues and limitations with Harmony's P2P layer

- The current P2P code is spread out among multiple packages (e.g. p2p, node, consensus) and there's no real separation of concerns in the code base.
- Testing and mocking is substandard compared to other chains - the bulk of Harmony's P2P code does not have unit tests or mocks.
- The current P2P code utilizes a custom categorization system to separate various message types by prefixing p2p messages with a set of bytes distinguishing the message category - this adds extra logic and code overhead compared to just using distinct topics.

### Separation of concerns
The current p2p package only contains ~400 lines of implementation code, which is quite different compared to how the majority of other blockchain projects structure their code.

Other projects have more obvious and encapsulated p2p packages which makes it:

- easier for developers to understand the underlying p2p implementation,
- easier to debug eventual p2p problems (since code is primarily encapsulated in one given package), and
- easier to implement a higher degree of code coverage through unit testing and mocking.

### Testing standards
Most other projects have code coverage of around ~40-60% of their p2p code, Harmony has around <~15% coverage (when factoring in p2p centric code in node and consensus packages).

Nil pointers/dereferences shouldn't happen in production grade code and it's still somewhat of a reoccurring event in the current code base. If the goal is to run financial and DeFi related apps on Harmony, we need to step up the testing and coverage efforts.

### Message categorization & topics
Most other projects use properly separated topics, wherein a given topic is solely focused on processing messages of a given type. Dedicated topic queues for a given message type would:

- Make the code a lot cleaner since serialization and deserialization can be greatly simplified (no need to prefix messages with their category).
- Allow for even more fine-grained concurrency control of every respective topic (semaphores with different weighting).

## Goals

- **Separation of concerns:** p2p communication should be carried out by the p2p package and not spread out in the node, consensus and p2p packages - this makes it easier to debug p2p code and for external developers to easily understand Harmony's p2p architecture
- **Testability and code coverage:** we have to ensure that code shipped to production doesn't contain run time bugs - these should be discovered by unit tests before deployment to production
- **Performance, scalability and attack mitigation:** by introducing separate pubsub topics per distinct message type we can also use tailored semaphore weights for every given message type and have full controll of message throughput to various components. If the serialization format is also changed from RLP to ProtoBuf, a lot of serialization and deserialization performance gains can also be made.

## Implementation
Implementation will be carried out in three separate phases and *the goal is to focus on small and incremental PR:s as much as possible*:

- **Phase 1:** Improving code coverage of existing node, consensus and p2p packages without extensive code refactoring (backwards compatible)
- **Phase 2:** Separation of concerns: Refactoring + Moving p2p related code from node & consensus packages to p2p package, implementing interface driven design for p2p package (backwards compatible)
- **Phase 3:** Topic / message categorization refactoring, fine-grained concurrency control, protobuf for serialization wherever possible (not backwards compatible)

### Phase 1: Improving code coverage of existing node, consensus and p2p packages without extensive code refactoring (backwards compatible)
Phase 1 revolves around drastically improving the code coverage for the node, consensus and existing p2p packages. I've already started some work on this phase, but there's a lot of work left to be done in order to get the coverage up to a similar code coverage standard of what competing projects currently have.

The unit tests will serve as an implementation contract and compatibility check for the ensuing phases. 

**Actions:**

- Create more unit tests / improve code coverage
- Minor code refactoring: keep existing structs, interfaces etc as much as possible - but break up massive function definitions into smaller and more easily tested functions

### Phase 2: Separation of concerns: Refactoring + Moving p2p related code from node & consensus packages to p2p package, implementing interface driven design for p2p package (backwards compatible)
Phase 2 revolves around separation of concerns and moving p2p centric code from the node and consensus packages over to the p2p package and refactoring the code to adhere to a set of interfaces (see the *"Appendix A: Interfaces / pseudo-code"* section below).

Using an interface driven design will allow for greater testability and mocking, improved quality assurance (through improved code coverage), and simultaneously allowing for greater implementation flexibility when resharding and fast sync is implemented.

**Actions:**

- Introduce relevant interfaces from "Appendix A: Interfaces / pseudo-code" supporting the current code base at the time of Phase 2
- Refactor implementation code and tests/mocks to adhere to the new set of interfaces

### Phase 3: Topic / message categorization refactoring, fine-grained concurrency control, protobuf for serialization wherever possible (not backwards compatible)
The last phase centers around moving towards fully separated message topic queues and fine-grained concurrency per given message type/category.

This means that instead of using a few generalized topics and parsing out message types/categories via the byte prefix, all distinct message types would use distinct message topics, e.g:

- `hmy/mainnet/0.0.1/shard/0/consensus // consensus messages for the beacon shard / shard 0`
- `hmy/mainnet/0.0.1/shard/1/consensus // consensus messages for shard 1`
- `hmy/mainnet/0.0.1/shard/0/transactions // transaction messages for shard 0`
- `hmy/mainnet/0.0.1/shard/0/1/transactions // transaction messages for 0 -> 1 cross shard`
- `hmy/mainnet/0.0.1/shard/1/crosslinks // crosslinks beacon -> shard 1`
- `hmy/mainnet/0.0.1/shard/2/crosslinks // crosslinks beacon -> shard 2`

Alternatively:

- `consensus_0 // consensus messages for the beacon shard / shard 0`
- `consensus_1 // consensus messages for shard 1`
- `transactions_0 // transaction messages for shard 0`
- `transactions_0_1 // transaction messages for 0 -> 1 cross shard`
- `crosslinks_1 // crosslinks beacon -> shard 1`
- `crosslinks_2 // crosslinks beacon -> shard 2`

**Benefits vs current design:**

- Semaphore weighting can be fine-tuned for each distinct message type/category - consensus messages & crosslinks can e.g. be prioritized over tx messages etc.
- Message byte prefixes can be removed which will lead to simpler and more maintainable serialization and deserialization code of p2p messages.

*NOTE: This change would break backwards compatibility - and if implemented, would require extensive coordination with external validators.*

#### Serialization & Deserialization
We could also explore the idea of moving more message serialization and deserialization over to Protobuf instead of RLP.

[RLP is a lot slower compared to Protobuf](https://github.com/prysmaticlabs/prysm/issues/139) and a lot of ETH 2 nodes & clients are [moving in the direction of dropping RLP](https://github.com/prysmaticlabs/prysm/issues/150) in favor of a faster serialization and deserialization protocol.

A lot of newer layer 1's don't actually use RLP at all, or are [in the process](https://github.com/Fantom-foundation/go-lachesis/issues/158) [of moving away from it](https://github.com/Fantom-foundation/go-lachesis/pull/163) since e.g. *Protobuf is at minimum 2x faster than RLP*.

## External Impact
Phase 1 and 2 would have minimal external impact on validators since they would be backwards compatible with current protocols.

Phase 3 would break backwards compatibility and would require extensive coordination with the validators.

**Phase 3 (if approved) should only go live after extremely thorough testing on LRTN and after detailed communication with the external validator community explaining the upcoming protocol change.**
