import axios from "axios";
import useSWR, { SWRConfiguration } from "swr";

export const baseURL = "/api/";

export const api = axios.create({
  baseURL,
});

export const useApi = <T>(url: string | null, config?: SWRConfiguration<T>) => {
  return useSWR<T, ApiError<ApiResponse<T>>>(
    url,
    (url: string) => api.get(url).then((res) => res.data.data),
    config,
  );
};

export let token = localStorage.getItem("token");
export const setToken = (t: string) => {
  token = t;
  localStorage.setItem("token", t);
};

export const goLogin = () => {
  location.assign("/login");
};

export function apiV1ErrorHandler(err: ApiError<void>) {
  switch (true) {
    case err.name === "CanceledError":
      break;
    case !err || !err.response || !err.response.data:
      err.msg = "网络错误";
      break;
    default:
      err.msg = err.response?.data?.msg;
  }
  return err;
}

api.interceptors.request.use((config) => {
  if (config.url?.startsWith("public/")) {
    return config;
  }
  if (!token) {
    goLogin();
    const controller = new AbortController();
    controller.abort();
    return {
      ...config,
      signal: controller.signal,
    };
  }
  config.headers.Authorization = token;
  return config;
});
api.interceptors.response.use(undefined, (err: ApiError<ApiResponse<void>>) => {
  console.log(err);
  if (err.response?.data?.code === 3) {
    goLogin();
    return new Promise(() => {});
  }
  apiV1ErrorHandler(err);
  return Promise.reject(err);
});
