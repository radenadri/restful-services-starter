import { integer, pgTable, serial, text, varchar } from 'drizzle-orm/pg-core';

import { users } from './users';

/**
 * @swagger
 * components:
 *   schemas:
 *     Post:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The ID of the posts.
 *           example: 1
 *         title:
 *           type: string
 *           description: The title of the posts.
 *           example: This is the title of the posts
 *         body:
 *           type: string
 *           description: The content of the posts
 *           example: Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s
 *         user_id:
 *           type: integer
 *           description: The user ID, the owner of the posts.
 *           example: 1
 */
export const posts = pgTable('posts', {
    id: serial('id').primaryKey(),
    title: varchar('title', { length: 50 }).notNull(),
    body: text('body'),
    userId: integer('user_id')
        .references(() => users.id)
        .notNull(),
});
