import { initContract } from "@ts-rest/core";
import { categoryContract } from "./category.js";
import { commentContract } from "./comment.js";
import { healthContract } from "./health.js";
import { todoContract } from "./todo.js";

const c = initContract();

export const apiContract = c.router({
   Health: healthContract,
  Todo: todoContract,
  Comment: commentContract,
  Category: categoryContract,
});
