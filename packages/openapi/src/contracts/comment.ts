import { ZTodoComment } from "@boilerplate/zod";
import { initContract } from "@ts-rest/core";
import z from "zod";
import { getSecurityMetadata } from "../utils.js";

const c = initContract();

const metadata = getSecurityMetadata();

export const commentContract = c.router(
  {
    addComment: {
      summary: "Add comment to todo",
      path: "/todos/:id/comments",
      method: "POST",
      body: ZTodoComment.pick({
        content: true,
      }),
      responses: {
        201: ZTodoComment,
      },
      metadata: metadata,
    },

    getCommentsByTodoId: {
      summary: "Get comments for todo",
      path: "/todos/:id/comments",
      method: "GET",
      responses: {
        200: z.array(ZTodoComment),
      },
      metadata: metadata,
    },

    updateComment: {
      summary: "Update comment",
      path: "/comments/:id",
      method: "PATCH",
      body: ZTodoComment.pick({
        content: true,
      }),
      responses: {
        200: ZTodoComment,
      },
      metadata: metadata,
    },

    deleteComment: {
      summary: "Delete comment",
      path: "/comments/:id",
      method: "DELETE",
      responses: {
        204: z.void(),
      },
      metadata: metadata,
    },
  },
  {
    pathPrefix: "/v1",
  }
);
