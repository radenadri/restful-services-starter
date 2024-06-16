/** @type {import('ts-jest').JestConfigWithTsJest} */
export default {
    clearMocks: true,
    coverageProvider: 'v8',
    moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx', 'json', 'node'],
    moduleNameMapper: {
        '@/(.*)': '<rootDir>/src/$1',
    },
    roots: ['<rootDir>/src'],
    testPathIgnorePatterns: ['/node_modules/'],
    testMatch: ['**/__tests__/**/*.ts?(x)', '**/?(*.)+(spec|test).ts?(x)'],
    transform: {
        '^.+\\.tsx?$': 'ts-jest',
    },
};
