import { drizzle } from 'drizzle-orm/postgres-js';
import postgres from 'postgres';

import { DB_DATABASE, DB_HOST, DB_MAX_CONNECTIONS, DB_PASSWORD, DB_PORT, DB_USER } from '../config';

const pg = postgres({
    host: DB_HOST,
    port: DB_PORT,
    user: DB_USER,
    password: DB_PASSWORD,
    database: DB_DATABASE,
    max: DB_MAX_CONNECTIONS,
});

const db = drizzle(pg);

/* auto migrate, uncomment to enable */
// migrate(db, { migrationsFolder: 'drizzle' });

export default db;
