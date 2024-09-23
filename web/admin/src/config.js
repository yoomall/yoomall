
const prod = {
    url: {
        API_URL: '/api/v1',
        BASE_URL: '/',
        IMG_URL: '/',
    }
}

const dev = {
    url: {
        API_URL: 'http://localhost:8900/api/v1',
        BASE_URL: 'http://localhost:8000',
        IMG_URL: 'http://localhost:8000/',
    }
}

const debug = import.meta.env.MODE === 'development'

export const config = debug ? dev : prod
export default config