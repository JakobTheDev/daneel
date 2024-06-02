CREATE TABLE [dbo].[Program]
(
  [Id] INT NOT NULL IDENTITY,
  [PlatformId] INT NOT NULL,
  [DisplayName] VARCHAR (50) NOT NULL,
  [IsPrivate] BIT DEFAULT 0,
  [IsActive] BIT DEFAULT 1,

  CONSTRAINT [PK_Program] PRIMARY KEY CLUSTERED ([Id] ASC),
  CONSTRAINT [FK_Program_PlatformId] FOREIGN KEY (PlatformId) REFERENCES Platform([Id])
)
