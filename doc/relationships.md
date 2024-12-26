Relationships in GORM
User ↔ Account

A User has many Account.
Each Account belongs to a User.
Account ↔ AccountInformation

An Account has many AccountInformation.
Each AccountInformation belongs to an Account.
Account ↔ Transaction

An Account can send or receive multiple Transactions.
Transaction ↔ ExternalBank

A Transaction optionally belongs to an ExternalBank.
