import { desc, eq, sql } from 'drizzle-orm';

import db from '@/db';
import { posts } from '@/models/posts';

export async function getAllPosts(page: number = 1, perPage: number = 10) {
    return await db
        .select()
        .from(posts)
        .orderBy(desc(posts.id))
        .limit(perPage)
        .offset((page - 1) * perPage);
}

export async function getPost(postId: number) {
    const res = await db.select().from(posts).where(eq(posts.id, postId)).limit(1);

    return res[0];
}

export async function getPostsByUser(userId: number) {
    return await db.select().from(posts).where(eq(posts.userId, userId));
}

export async function createPost(title: string, body: string, userId: number) {
    return await db.insert(posts).values({ title, body, userId }).returning();
}

export async function updatePost(body: string, userId: number, postId: number) {
    const res = await db
        .update(posts)
        .set({ body })
        .where(sql`${posts.userId} = ${userId} and ${posts.id} = ${postId}`)
        .returning();

    return res[0];
}

export async function deletePost(postId: number, userId: number) {
    await db.delete(posts).where(sql`${posts.userId} = ${userId} and ${posts.id} = ${postId}`);
}
