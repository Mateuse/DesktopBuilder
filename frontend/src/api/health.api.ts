// Checks health of backend services
import { ROUTES } from "../constants/urls"


// Golang main backend
export const mainHealth = async () => {
    const response = await fetch(ROUTES.MAIN_BE_HEALTH)
    if (!response.ok) {
        throw new Error("Failed to fetch health")
    }
    const data = await response.json()
    return data;
}