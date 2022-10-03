import config from './config';
// import jwtDecode from 'jwt-decode';
// import * as moment from 'moment';

const axios = require('axios');


class FiberClient {
  constructor(overrides) {
    this.config = {
      ...config,
      ...overrides,
    };
    this.authToken = config.authToken;

    this.login = this.login.bind(this);
    this.apiClient = this.getApiClient(this.config);
  }

  /* ----- Authentication & User Operations ----- */

  /* Authenticate the user with the backend services.
	 * The same JWT should be valid for both the api and cms */
  login(email, password) {
    //delete this.apiClient.defaults.headers['Authorization'];
    
    // HACK: This is a hack for scenario where there is no login form
    const form_data = new FormData();
    const grant_type = 'password';
    const item = {grant_type, email, password};
    for (const key in item) {
      form_data.append(key, item[key]);
    }
    console.log("sugma");
    // return this.apiClient
    //     .post('/auth/login', form_data)
    //     .then((resp) => {
    //       localStorage.setItem('token', JSON.stringify(resp.data));
    //       return this.fetchUser();
    //     });
    return this.apiClient
    .post('http://localhost:8000/auth/login', {
        email: email,
        password: password,
      })
    .then((resp) => {
      localStorage.setItem('token', resp.data.token);
      localStorage.setItem('email', email);

      //return this.fetchUser();
    });

  }

  fetchUser() {
    const formData = new FormData();
    formData.append('email', localStorage.getItem('email'));
    const config = {
      params:{email: localStorage.getItem('email')},
      headers: {Authorization:`Bearer ${localStorage.getItem('token')}` }
    };

    console.log(JSON.stringify(config))
    console.log(JSON.stringify(localStorage.getItem('data')))
    return this.apiClient.get('http://localhost:8000/users', config
    ).then(({data}) => {
      localStorage.setItem('user', JSON.stringify(data[0]));
      return data;
    });
  }

  register(email, password, fullName) {
    const registerData = {
      email,
      password,
      full_name: fullName,
      is_active: true,
    };

    return this.apiClient.post('/auth/signup', registerData).then(
        (resp) => {
          return resp.data;
        });
  }

  // Logging out is just deleting the jwt.
  logout() {
    // Add here any other data that needs to be deleted from local storage
    // on logout
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  }

  /* ----- Client Configuration ----- */

  /* Create Axios client instance pointing at the REST api backend */
  getApiClient(config) {
    const initialConfig = {
      //baseURL: `${config.apiBasePath}/api/v1`,
      baseURL: `${config.apiBasePath}`,
    };
    const client = axios.create(initialConfig);
    //client.interceptors.request.use(localStorageTokenInterceptor);
    return client;
  }

  getModel3D(model3dId) {
    return this.apiClient.get(`/model3ds/${model3dId}`);
  }

  getModel3Ds(keyword) {
    return this.apiClient.get(`/model3ds/search/?keyword=${keyword}&max_results=10`).then(({data}) => {
      return data;
    });
  }

  getUserModel3Ds() {
    return this.apiClient.get(`/model3ds/my-model3ds/`).then(({data}) => {
      return data;
    });
  }

  createModel3D(title, author, description, price, blob_data, file_name) {
    // console.log("author: "+author)
    // console.log("file_name: " + file_name)
    console.log(localStorage.getItem('user'))
    const model3dData = {
      Title : title,
      Author : localStorage.getItem('user')["user_name"], 
      Description : description, 
      Price : Math.round(price * 100) / 100,
      BlobData : blob_data, 
      FileName : file_name,
      // slocalStorage.getItem('user')bmitter_id: submitter_id,
    };
    const config = {
      headers: {Authorization:`Bearer ${localStorage.getItem('token')}` }
    }
    return this.apiClient.post(`http://localhost:8000/model3ds`, model3dData, config);
  }


  deleteModel3D(model3dId) {
    return this.apiClient.delete(`/model3ds/${model3dId}`);
  }
}


// every request is intercepted and has auth header injected.
// function localStorageTokenInterceptor(config) {
//   const headers = {};
//   const tokenString = localStorage.getItem('token');

//   if (tokenString) {
//     const token = JSON.parse(tokenString);
//     const decodedAccessToken = jwtDecode(token.token);
//     const isAccessTokenValid =
// 			moment.unix(decodedAccessToken.exp).toDate() > new Date();
//     if (isAccessTokenValid) {
//       headers['Authorization'] = `Bearer ${token.token}`;
//     } else {
//       alert('Your login session has expired');
//     }
//   }
//   config['headers'] = headers;
//   return config;
// }

export default FiberClient;
