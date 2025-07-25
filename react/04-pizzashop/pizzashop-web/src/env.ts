import z4 from "zod/v4";

const envSchema = z4.object({ VITE_API_URL: z4.url() });

export const env = envSchema.parse(import.meta.env);
