export const MAIN_BE_URL = import.meta.env.VITE_MAIN_BACKEND_API || "http://localhost:8080"


export const ROUTES = {
    MAIN_BE_HEALTH: `${MAIN_BE_URL}/health`,
}