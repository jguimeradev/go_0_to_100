# env-reader

## Example: A Web API Service

Imagine you're building a simple web API service that needs to connect to a database and expose some endpoints.

### Configuration File (config.yaml)

Your config file might contain these **default settings**:
- Server port: 8080
- Database host: localhost
- Database port: 5432
- Database name: myapp_db
- Log level: info
- Request timeout: 30 seconds
- Max connections to database: 10

These are your "baseline" settings that work for local development.

### Environment Variables

Now when you deploy to **production**, you don't want to change the config file. Instead, you set environment variables:
- `DB_HOST=prod-database.company.com` (overrides localhost)
- `DB_PASSWORD=secretpassword123` (sensitive data, never in config files!)
- `LOG_LEVEL=error` (less verbose in production)
- `PORT=3000` (maybe production uses a different port)

### How Your Program Would Use This

When your program starts:

1. It first loads the config file and gets all the default values
2. Then it checks for environment variables
3. If an env var exists for a setting, it **overrides** the config file value
4. Sensitive data like passwords **only** come from env vars, never stored in files

### Final Result

Your running application would use:
- Port: **3000** (from env var, not 8080)
- Database host: **prod-database.company.com** (from env var, not localhost)
- Database port: **5432** (from config file, no override)
- Database password: **secretpassword123** (from env var only)
- Log level: **error** (from env var, not info)
- Timeout: **30 seconds** (from config file, no override)

### Why This Pattern?

- Config file can be committed to version control safely (no secrets)
- Same code/config file works in dev, staging, and production
- Each environment just sets different env vars
- Secrets never accidentally get committed to Git
- Easy to change production settings without redeploying code

Does this example make the exercise clearer?