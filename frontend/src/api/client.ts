import axios from 'axios';

const apiClient = axios.create({
    baseURL: process.env.REACT_APP_API_BASE_URL || 'http://localhost:8080/',
    headers: {
        'Content-Type': 'application/json',
    },
});

apiClient.interceptors.response.use(
    response => response,
    error => {
        // Правильное извлечение сообщения об ошибке
        const message = error.response?.data?.message || error.message;
        const statusCode = error.response?.status || 500;

        return Promise.reject({
            message: statusCode === 409
                ? `Team already exists: ${message}`
                : message,
            statusCode
        });
    }
);

export default apiClient;