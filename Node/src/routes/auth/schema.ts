import { z } from 'zod';

export const loginSchema = z.object({
    body: z.object({
        email: z.string().min(1, 'Please enter your email').email(),
        password: z.string().min(1, 'Please enter your password'),
    }),
});

export const registerSchema = z.object({
    body: z.object({
        name: z.string().min(1, 'Please enter your name'),
        email: z.string().min(1, 'Please enter your email').email(),
        password: z.string().min(8, 'Password must be at least 8 characters'),
        password_confirmation: z.string().min(8, 'Please confirm your password'),
    }),
});

export const generateTokenSchema = z.object({
    body: z.object({
        id: z.number(),
        email: z.string().min(1, 'Please enter your email').email(),
        name: z.string().min(1, 'Please enter your name'),
    }),
});
