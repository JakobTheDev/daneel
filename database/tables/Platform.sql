CREATE TABLE [dbo].[Platform]
(
  [Id] INT NOT NULL IDENTITY,
  [Name] VARCHAR (50) NOT NULL,
  [IsActive] BIT DEFAULT 1,

  CONSTRAINT [PK_Platform] PRIMARY KEY CLUSTERED ([Id] ASC)
)
