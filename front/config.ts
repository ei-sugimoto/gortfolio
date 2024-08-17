type Config = {
  url: string;
};

const config: Config = {
  url: process.env.API_URL ?? "",
};

if (!config.url) {
  throw new Error("API_URL is not defined");
}

export default config;
