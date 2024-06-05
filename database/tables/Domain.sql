CREATE TABLE [dbo].[Domain]
(
  [Id] INT NOT NULL IDENTITY,
  [ProgramId] INT NOT NULL,
  [DomainName] VARCHAR (256) NOT NULL,
  [IsInScope] BIT DEFAULT 1,
  [IsActive] BIT DEFAULT 1,

  CONSTRAINT [PK_Domain] PRIMARY KEY CLUSTERED ([Id] ASC),
  CONSTRAINT [FK_Domain_ProgramId] FOREIGN KEY (ProgramId) REFERENCES Program([Id])
)
