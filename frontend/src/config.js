import runtimeEnv from '@mars/heroku-js-runtime-env';

const env = runtimeEnv();
const config = {
  apiBasePath: env.REACT_APP_API_BASE_PATH || `http://localhost:8000`,//'https://fastapi-model3d-app.herokuapp.com',
  reactAppMode: process.env.REACT_APP_MODE || 'dev',
};

export default config;
