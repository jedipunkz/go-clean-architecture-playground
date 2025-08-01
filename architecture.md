# Go クリーンアーキテクチャ構成図

## システム全体構成図

```mermaid
graph TB
    subgraph "外部"
        Client[HTTPクライアント]
    end
    
    subgraph "Interface Layer（インターフェース層）"
        Controller[UserController<br/>HTTP Handler]
        RepoInterface[UserRepository<br/>Interface]
    end
    
    subgraph "Use Case Layer（ユースケース層）"
        UseCase[UserUsecase<br/>ビジネスロジック]
    end
    
    subgraph "Entity Layer（エンティティ層）"
        Entity[User Entity<br/>ビジネスルール]
    end
    
    subgraph "Infrastructure Layer（インフラ層）"
        Repository[MemoryUserRepository<br/>データ永続化]
    end
    
    subgraph "Main（依存性注入）"
        Main[main.go<br/>DI Container]
    end

    Client --> Controller
    Controller --> UseCase
    UseCase --> RepoInterface
    UseCase --> Entity
    Repository -.-> RepoInterface
    Main --> Controller
    Main --> UseCase
    Main --> Repository

    classDef entity fill:#e1f5fe,stroke:#0277bd,stroke-width:2px
    classDef usecase fill:#f3e5f5,stroke:#7b1fa2,stroke-width:2px
    classDef interface fill:#e8f5e8,stroke:#2e7d32,stroke-width:2px
    classDef infrastructure fill:#fff3e0,stroke:#f57c00,stroke-width:2px
    classDef main fill:#fce4ec,stroke:#c2185b,stroke-width:2px

    class Entity entity
    class UseCase usecase
    class Controller,RepoInterface interface
    class Repository infrastructure
    class Main main
```

## 依存関係の方向

```mermaid
graph LR
    Infrastructure[Infrastructure<br/>インフラ層] --> Interface[Interface<br/>インターフェース層]
    Interface --> UseCase[Use Case<br/>ユースケース層]
    UseCase --> Entity[Entity<br/>エンティティ層]
    Main[Main<br/>DI] --> Infrastructure
    Main --> Interface
    Main --> UseCase

    classDef layer fill:#f9f9f9,stroke:#333,stroke-width:2px
    class Infrastructure,Interface,UseCase,Entity,Main layer
```

## データフロー（API リクエスト処理）

```mermaid
sequenceDiagram
    participant C as Client
    participant Ctrl as UserController
    participant UC as UserUsecase
    participant E as User Entity
    participant R as UserRepository

    C->>Ctrl: POST /users
    Ctrl->>UC: CreateUser(name, email)
    UC->>E: NewUser(name, email)
    E-->>UC: User entity
    UC->>R: Create(user)
    R-->>UC: Success
    UC-->>Ctrl: Created user
    Ctrl-->>C: JSON response
```

## レイヤー構成の詳細

```mermaid
graph TD
    subgraph "Clean Architecture Layers"
        subgraph "Entity Layer"
            E1[User Entity<br/>- ID, Name, Email<br/>- NewUser()<br/>- UpdateInfo()]
        end
        
        subgraph "Use Case Layer"
            U1[UserUsecase<br/>- CreateUser()<br/>- GetUser()<br/>- UpdateUser()<br/>- DeleteUser()]
        end
        
        subgraph "Interface Layer"
            I1[UserController<br/>- HTTP Handlers]
            I2[UserRepository<br/>- Interface Definition]
        end
        
        subgraph "Infrastructure Layer"
            F1[MemoryUserRepository<br/>- Concrete Implementation<br/>- In-memory storage]
        end
    end

    U1 --> E1
    U1 --> I2
    I1 --> U1
    F1 -.-> I2

    classDef entity fill:#e1f5fe
    classDef usecase fill:#f3e5f5
    classDef interface fill:#e8f5e8
    classDef infrastructure fill:#fff3e0

    class E1 entity
    class U1 usecase
    class I1,I2 interface
    class F1 infrastructure
```