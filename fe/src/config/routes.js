import { config } from "./config";

config.routes = {
  auth: {
    base: "/auth",
    login: "/login",
  },

  expenses: {
    base: "/expenses",
    approve: "/approve",
    reject: "/reject",
  },
};
