CREATE TABLE [dbo].[Subdomain]
(
  [Id] INT NOT NULL IDENTITY,
  [DomainId] INT NOT NULL,
  [SubdomainName] VARCHAR (256) NOT NULL,
  [IsInScope] BIT DEFAULT 1,
  [IsActive] BIT DEFAULT 1,

  CONSTRAINT [PK_Subdomain] PRIMARY KEY CLUSTERED ([Id] ASC),
  CONSTRAINT [FK_Subdomain_DomainId] FOREIGN KEY (DomainId) REFERENCES Domain([Id])
)
