import z4 from "zod/v4";

const envSchema = z4.object({
  VITE_API_URL: z4.url(),
  VITE_ENABLE_API_DELAY: z4.string().transform((value) => value === "true"),
});

export const env = envSchema.parse(import.meta.env);
