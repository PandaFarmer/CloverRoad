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
  // const [userName, setUserName] = useState("");
  /* ----- Authentication & User Operations ----- */
  
  /* Authenticate the user with the backend services.
	 * The same JWT should be valid for both the api and cms */
  login(email, password) {
    //delete this.apiClient.defaults.headers['Authorization'];
    
    // HACK: This is a hack for scenario where there is no login form
    // const form_data = new FormData();
    // const grant_type = 'password';
    // const item = {grant_type, email, password};
    // for (const key in item) {
    //   form_data.append(key, item[key]);
    // }
    console.log("attempting to send request along proper route");
    return this.apiClient
    .post("auth/login",{
      // headers: {
      //   'Access-Control-Allow-Origin': true,
      // },
        email: email,
        password: password,
      })
    .then((resp) => {
      localStorage.setItem('token', resp.data.token);
      localStorage.setItem('email', email);
      console.log("fetching user................")
      return this.fetchUser();
    });

  }

  register(email, username, fullname, password) {

    return this.apiClient.post('/auth/signup', {
      email : email,
      username : username,
      fullname : fullname,
      password : password
    }).then(
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

  getModel3Ds(keywords, max_results, lastUpdated, minPrice, maxPrice) {
    console.log(lastUpdated.toString());
    const data = new FormData();
    data.append("keywords", keywords);
    data.append("max_results", max_results);
    data.append("lastUpdated", lastUpdated);
    data.append("minPrice", minPrice);
    data.append("maxPrice", maxPrice);

    const config = {
      data,
      headers: {Authorization:`Bearer ${localStorage.getItem('token')}` , 
        // 'Access-Control-Allow-Origin': '*',
        // 'Content-Type': 'multipart/form-data'
        'Content-Type':'application/json'
      }
    }

    console.log(config);
    return this.apiClient.get(`/model3ds/search`, config).then(({data}) => {
      return data;
    });
  }

  getUserModel3Ds() {
    const config = {
      params:{user: localStorage.getItem('user')},
      headers: {Authorization:`Bearer ${localStorage.getItem('token')}` }
    };

    return this.apiClient.get(`/model3ds`, config).then(({data}) => {
      console.log(data);
      return data;
    });
  }

  createModel3D(title, author, description, price, previewBlobData, blobData, fileName) {
    // console.log("author: "+author)
    // console.log("file_name: " + file_name)
    console.log("user in localStorage: " + localStorage.getItem('user'));
    const data = new FormData();
    data.append("title", title);
    data.append("author", localStorage.getItem('user'));
    data.append("description", description);
    data.append("price", Math.floor(price * 100) / 100);
    // data.append("serialized_preview_file", previewBlobData)
    data.append("serialized_file_3d", blobData);
    data.append("file_name_and_extension", fileName);

    const config = {
      headers: {Authorization:`Bearer ${localStorage.getItem('token')}` , 
        // 'Access-Control-Allow-Origin': '*',                        
        // 'Content-Type': 'multipart/form-data',
        'Content-Type':'application/json'
      },
    };
    return this.apiClient.post(`/model3ds`, data, config
      ).catch(error => {
        console.log('Error: ', error)
      });
  }

  fetchUser() {
    console.log("qasdf;lkjasdfl;kjaf;'-cd")
    // const formData = new FormData();
    // formData.append('email', localStorage.getItem('email'));
    const config = {
      params:{email: localStorage.getItem('email')},
      headers: {//'Access-Control-Allow-Origin': true, 
        Authorization:`Bearer ${localStorage.getItem('token')}`
      }
    };
    console.log("config: " + config)
    console.log(JSON.stringify(config));
    console.log(JSON.stringify(localStorage.getItem('data')));
    console.log("baseurl: " + this.baseURL);
    return this.apiClient.get('users', config
    ).then(({data}) => {
      console.log("user: " + JSON.stringify(data[0]));
      localStorage.setItem('user', data[0].user_name);
      return data;
    });
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
