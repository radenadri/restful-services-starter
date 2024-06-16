import cookie from 'cookie';
import { Response } from 'express';
import jwt from 'jsonwebtoken';

import { JWT_SECRET_KEY } from '@/config';
import { UserData } from '@/interfaces/user-data';

export function createAccessToken(id: number, email: string, name: string) {
    const payload = { id, email, name };
    const token = jwt.sign(payload, JWT_SECRET_KEY);
    return token;
}

export function createRefreshToken(id: number, email: string, name: string) {
    const payload = { id, email, name };
    const token = jwt.sign(payload, JWT_SECRET_KEY);
    return token;
}

export function verifyToken(token: string) {
    try {
        return jwt.verify(token, JWT_SECRET_KEY) as UserData;
    } catch (error) {
        if (error instanceof jwt.TokenExpiredError) {
            throw new Error('Token expired');
        }

        throw new Error('Invalid token');
    }
}

export function setRefreshCookie(res: Response, refreshToken: string) {
    const date = new Date();
    date.setFullYear(date.getFullYear() + 1);
    res.setHeader(
        'Set-Cookie',
        cookie.serialize('refreshToken', refreshToken, {
            httpOnly: true,
            expires: date,
            sameSite: 'none',
            secure: true,
            path: '/',
        }),
    );
}
