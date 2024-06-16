import { NextFunction, Request, Response } from 'express';

import { UserData } from '@/interfaces/user-data';
import { verifyToken } from '@/routes/auth/utils';

declare module 'express-serve-static-core' {
    interface Request {
        user: { id: number; email: string; name: string };
    }
}

export default function authenticateUser(req: Request, res: Response, next: NextFunction) {
    const token = req.headers.authorization?.split(' ')[1];

    if (!token) {
        return res.status(401).send({
            success: false,
            error: 'Unauthorized',
        });
    }

    try {
        const { id, email, name }: UserData = verifyToken(token);

        req.user = { id, email, name };

        next();
    } catch (error: Error | unknown) {
        res.status(401).send({
            success: false,
            error: error instanceof Error ? error.message : 'Unauthorized',
        });
    }
}
