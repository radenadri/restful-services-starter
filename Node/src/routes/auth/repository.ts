import bcrypt from 'bcrypt';
import { eq } from 'drizzle-orm';

import db from '@/db';
import { users } from '@/models/users';

export async function findByEmail(email: string) {
    const [res] = await db.select().from(users).where(eq(users.email, email)).limit(1);

    return res;
}

export async function verify(email: string, password: string) {
    const user = await findByEmail(email);

    if (!user) return null;

    const passwordIsValid = await bcrypt.compare(password.replace(/^\$2y/, '$2b'), user.password);
    return passwordIsValid ? user : null;
}

export async function register(name: string, email: string, password: string, passwordConfirm: string) {
    // check if passwords match
    if (password !== passwordConfirm) {
        throw new Error('Passwords do not match');
    }

    // hash password
    const hashedPassword = await bcrypt.hash(password, 10);

    // get username from email before @
    const username = email.split('@')[0];

    // insert user
    const user = {
        name,
        username,
        email,
        password: hashedPassword.replace(/^\$2b/, '$2y'),
    };

    return await db
        .insert(users)
        .values(user)
        .returning({
            id: users.id,
            name: users.name,
            email: users.email,
        })
        .execute();
}
