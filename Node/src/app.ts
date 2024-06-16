import bodyParser from 'body-parser';
import cookies from 'cookie-parser';
import cors from 'cors';
import express from 'express';
import * as fs from 'fs';
import helmet from 'helmet';
import morgan from 'morgan';
import path from 'path';

import { API_VERSION, ORIGIN } from './config';
import errorHandler from './middlewares/error-handler';
import { rateLimiter } from './middlewares/rate-limiter';
import { specs, swaggerUi } from './middlewares/swagger';
import routes from './routes';

const app = express();

// create logs folder if it doesn't exist
if (!fs.existsSync(path.join(__dirname, 'logs'))) {
    fs.mkdirSync(path.join(__dirname, 'logs'));
}

const logFile = fs.createWriteStream(path.join(__dirname, 'logs', 'access.log'), { flags: 'a' });

/* Library Middlewares */
app.use(morgan('combined', { stream: logFile }));
app.use(helmet());
app.use(express.json());
app.use(cookies());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(
    cors({
        credentials: true,
        origin: ORIGIN,
        methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    }),
);

/* App Middlewares */
app.use(errorHandler);
app.use(rateLimiter);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(specs));

/* Create public directory if it doesn't exist */
if (!fs.existsSync(path.join(__dirname, 'public'))) {
    fs.mkdirSync(path.join(__dirname, 'public'));
}

/* Include assets & public directory */
app.use(express.static(path.join(__dirname, 'assets')));
app.use(express.static(path.join(__dirname, 'public')));

/* Routing */
app.get('/', (_req, res) => res.send('It works!'));
app.use(`/api/${API_VERSION}`, routes);

export default app;
