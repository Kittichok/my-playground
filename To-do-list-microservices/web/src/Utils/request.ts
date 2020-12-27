import axios, { AxiosRequestConfig, Method } from 'axios';

//TODO mock to call api
export async function request(
  url: string,
  method: Method,
  options?: RequestInit
): Promise<{} | { err: any }> {
  const config: AxiosRequestConfig = {
    url,
    method,
  };
  if (options) {
    config.data = options;
  }
  const req = await axios(config);
  return req;
}
