-- This file contains SQL statements that will be executed after the build script.

-- ad db user
CREATE USER [daneel] FOR LOGIN [daneel] WITH DEFAULT_SCHEMA=[dbo]
GO
ALTER ROLE [db_owner] ADD MEMBER [daneel]
GO
