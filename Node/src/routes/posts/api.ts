import express from 'express';

import authenticateUser from '@/middlewares/authenticate-user';
import validate from '@/middlewares/validate-request';

import { createPost, deletePost, getAllPosts, getPost, getPostsByUser, updatePost } from './repository';
import { createPostSchema, updatePostSchema } from './schema';

const router = express.Router();

/**
 * @swagger
 * components:
 *   schemas:
 *     Post:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The post ID.
 *           example: 1
 *         title:
 *           type: string
 *           description: The post title.
 *           example: Hello World
 *         body:
 *           type: string
 *           description: The post body.
 *           example: This is my first post
 *         user_id:
 *           type: integer
 *           description: The user ID.
 *           example: 1
 *
 * /posts:
 *   get:
 *     summary: Get all posts
 *     description: Get all posts and return them as a json
 *     tags: [Post]
 *     parameters:
 *       - in: header
 *         name: Content-Type
 *         schema:
 *           type: string
 *           default: application/json
 *         required: true
 *       - in: query
 *         name: page
 *         schema:
 *           type: string
 *           default: 1
 *         required: false
 *         description: The page number to access
 *       - in: query
 *         name: perPage
 *         schema:
 *           type: string
 *           default: 10
 *         required: false
 *         description: The number of posts taken per page
 *     responses:
 *       200:
 *           description: Successfully get all posts
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       posts:
 *                         type: array
 *                         items:
 *                           $ref: '#/components/schemas/Post'
 *       400:
 *           description: Something went wrong getting your posts
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Something went wrong getting your posts
 */
router.get(
    '/',
    async (
        req: {
            query: {
                page: string;
                perPage: string;
            };
        },
        res,
    ) => {
        try {
            const page: number = req.query.page ? parseInt(req.query.page) : 1;
            const perPage: number = req.query.perPage ? parseInt(req.query.perPage) : 10;

            const posts = await getAllPosts(page, perPage);
            return res.status(200).send({ posts });
        } catch (error) {
            return res.status(400).send({ error: 'Something went wrong getting your posts' });
        }
    },
);

/**
 * @swagger
 * components:
 *   schemas:
 *     Post:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The post ID.
 *           example: 1
 *         title:
 *           type: string
 *           description: The post title.
 *           example: Hello World
 *         body:
 *           type: string
 *           description: The post body.
 *           example: This is my first post
 *         user_id:
 *           type: integer
 *           description: The user ID.
 *           example: 1
 *
 * /posts/users/{userId}:
 *   get:
 *     summary: Get all posts by user
 *     description: Get all posts by user and return them as a json
 *     tags: [Post]
 *     parameters:
 *       - in: header
 *         name: Content-Type
 *         schema:
 *           type: string
 *           default: application/json
 *         required: true
 *       - in: path
 *         name: userId
 *         schema:
 *           type: integer
 *         required: true
 *     responses:
 *       200:
 *           description: Successfully get all posts
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       posts:
 *                         type: array
 *                         items:
 *                           $ref: '#/components/schemas/Post'
 *       400:
 *           description: Something went wrong getting your posts
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Something went wrong getting your posts
 */
router.get('/users/:userId', async (req, res) => {
    try {
        const userId = parseInt(req.params.userId);
        const posts = await getPostsByUser(userId);
        res.status(200).send({ posts });
    } catch (error) {
        res.status(400).send({ error: 'Something went wrong getting posts' });
    }
});

/**
 * @swagger
 * components:
 *   securitySchemes:
 *     bearerAuth:
 *       type: http
 *       scheme: bearer
 *   schemas:
 *     Post:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The post ID.
 *           example: 1
 *         title:
 *           type: string
 *           description: The post title.
 *           example: Hello World
 *         body:
 *           type: string
 *           description: The post body.
 *           example: This is my first post
 *         user_id:
 *           type: integer
 *           description: The user ID.
 *           example: 1
 *
 * /posts:
 *   post:
 *     summary: Create a new post
 *     description: Create a new post and return it
 *     tags: [Post]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: header
 *         name: Content-Type
 *         schema:
 *           type: string
 *           default: application/json
 *         required: true
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               title:
 *                 type: string
 *                 description: The post title.
 *                 example: Hello World
 *               body:
 *                 type: string
 *                 description: The post body.
 *                 example: This is my first post
 *     responses:
 *       200:
 *           description: Successfully created post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       post:
 *                         type: object
 *                         $ref: '#/components/schemas/Post'
 *       400:
 *           description: Something went wrong creating your post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Something went wrong creating your post
 */
router.post('/', validate(createPostSchema), authenticateUser, async (req, res) => {
    try {
        const { title, body } = req.body;
        const userId = req.user.id;
        const post = await createPost(title, body, userId);
        return res.status(200).send({ post });
    } catch (error) {
        return res.status(400).send({ errror: 'Something went wrong creating your post' });
    }
});

/**
 * @swagger
 * components:
 *   schemas:
 *     Post:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The post ID.
 *           example: 1
 *         title:
 *           type: string
 *           description: The post title.
 *           example: Hello World
 *         body:
 *           type: string
 *           description: The post body.
 *           example: This is my first post
 *         user_id:
 *           type: integer
 *           description: The user ID.
 *           example: 1
 *
 * /posts/{postId}:
 *   get:
 *     summary: Get post by ID
 *     description: Get post by ID and return it as a json
 *     tags: [Post]
 *     parameters:
 *       - in: header
 *         name: Content-Type
 *         schema:
 *           type: string
 *           default: application/json
 *         required: true
 *       - in: path
 *         name: postId
 *         schema:
 *           type: integer
 *         required: true
 *     responses:
 *       200:
 *           description: Successfully get post by ID
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       post:
 *                         type: object
 *                         $ref: '#/components/schemas/Post'
 *       400:
 *           description: Something went wrong getting your post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Something went wrong getting your post
 *       404:
 *           description: Could not find post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Could not find post
 */
router.get('/:postId', async (req, res) => {
    try {
        const postId = parseInt(req.params.postId);
        const post = await getPost(postId);
        if (!post) return res.status(404).send({ error: 'Could not find post' });
        return res.status(200).send({ post });
    } catch (error) {
        return res.status(400).send({ error: 'Something went wrong getting your post' });
    }
});

/**
 * @swagger
 * components:
 *   securitySchemes:
 *     bearerAuth:
 *       type: http
 *       scheme: bearer
 *   schemas:
 *     Post:
 *       type: object
 *       properties:
 *         id:
 *           type: integer
 *           description: The post ID.
 *           example: 1
 *         title:
 *           type: string
 *           description: The post title.
 *           example: Hello World
 *         body:
 *           type: string
 *           description: The post body.
 *           example: This is my first post
 *         user_id:
 *           type: integer
 *           description: The user ID.
 *           example: 1
 *
 * /posts/{postId}:
 *   put:
 *     summary: Update post
 *     description: Update post and return it
 *     tags: [Post]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: header
 *         name: Content-Type
 *         schema:
 *           type: string
 *           default: application/json
 *         required: true
 *       - in: path
 *         name: postId
 *         schema:
 *           type: integer
 *         required: true
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               title:
 *                 type: string
 *                 description: The post title.
 *                 example: Hello World
 *               body:
 *                 type: string
 *                 description: The post body.
 *                 example: This is my first post
 *     responses:
 *       200:
 *           description: Successfully created post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       post:
 *                         type: object
 *                         $ref: '#/components/schemas/Post'
 *       400:
 *           description: Something went wrong updating your post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Something went wrong updating your post
 *       404:
 *           description: Could not find post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Could not find post
 */
router.put('/:postId', validate(updatePostSchema), authenticateUser, async (req, res) => {
    try {
        const { body } = req.body;
        const postId = parseInt(req.params.postId);
        const userId = req.user.id;

        const post = await getPost(postId);
        if (!post) return res.status(404).send({ error: 'Could not find post' });

        const updatedPost = await updatePost(body, userId, postId);
        res.status(200).send({ post: updatedPost });
    } catch (error) {
        res.status(400).send({ error: 'Something went wrong updating your post' });
    }
});

/**
 * @swagger
 * components:
 *   securitySchemes:
 *     bearerAuth:
 *       type: http
 *       scheme: bearer
 *
 * /posts/{postId}:
 *   delete:
 *     summary: Delete post
 *     description: Delete post and return the success deletion message
 *     tags: [Post]
 *     security:
 *       - bearerAuth: []
 *     parameters:
 *       - in: header
 *         name: Content-Type
 *         schema:
 *           type: string
 *           default: application/json
 *         required: true
 *       - in: path
 *         name: postId
 *         schema:
 *           type: integer
 *         required: true
 *     responses:
 *       200:
 *           description: Successfully deleted post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       message:
 *                         type: string
 *                         description: The success deletion message
 *                         example: Successfully deleted post
 *       400:
 *           description: Something went wrong deleting your post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Something went wrong deleting your post
 *       404:
 *           description: Could not find post
 *           content:
 *               application/json:
 *                   schema:
 *                     type: object
 *                     properties:
 *                       error:
 *                         type: string
 *                         description: The error message
 *                         example: Could not find post
 */
router.delete('/:postId', authenticateUser, async (req, res) => {
    try {
        const postId = parseInt(req.params.postId);
        const userId = req.user.id;

        const post = await getPost(postId);
        if (!post) return res.status(404).send({ error: 'Could not find post' });

        await deletePost(postId, userId);
        res.status(200).send({ message: 'Successfully deleted post' });
    } catch (error) {
        res.status(400).send({ error: 'Something went wrong deleting your post' });
    }
});

export default router;
