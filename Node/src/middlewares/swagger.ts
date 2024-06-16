import swaggerJsdoc from 'swagger-jsdoc';
import swaggerUi from 'swagger-ui-express';

import { API_VERSION, APP_URL } from '@/config';

const definition = {
    definition: {
        openapi: '3.1.0',
        info: {
            title: 'Express Boilerplate',
            version: '1.0.0',
            description: 'Express Boilerplate API',
            license: {
                name: 'MIT',
                url: 'https://spdx.org/licenses/MIT.html',
            },
            contact: {
                name: 'Adriana Eka Prayudha',
                url: 'https://radenadri.xyz',
                email: 'radenadriep@gmail.com',
            },
        },
        servers: [
            {
                url: `${APP_URL}/api/${API_VERSION}`,
            },
        ],
    },
    apis: ['**/*.ts'],
};

const specs = swaggerJsdoc(definition);
export { specs, swaggerUi };
