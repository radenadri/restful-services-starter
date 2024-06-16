import request from 'supertest';

import app from '../app';

describe('GET /', () => {
    it('should return 200', () => {
        request(app)
            .get('/')
            .expect(200)
            .end(err => {
                if (err) {
                    throw err;
                }
            });
    });

    it('should return 404', () => {
        request(app)
            .get('/404')
            .expect(404)
            .end(err => {
                if (err) {
                    throw err;
                }
            });
    });
});
