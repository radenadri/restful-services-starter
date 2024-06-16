import { pgTable, serial, varchar } from 'drizzle-orm/pg-core';

/**
 * @swagger
 * components:
 *   schemas:
 *     User:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The user ID.
 *           example: 1
 *         name:
 *           type: string
 *           description: The user name.
 *           example: Abraham Wong
 *         email:
 *           type: string
 *           description: The user email.
 *           example: wong@example.com
 *         password:
 *           type: string
 *           description: The user password.
 *           example: 123456
 */
export const users = pgTable('users', {
    id: serial('id').primaryKey(),
    name: varchar('name', { length: 50 }).notNull(),
    email: varchar('email', { length: 50 }).notNull().unique(),
    password: varchar('password', { length: 100 }).notNull(),
});
