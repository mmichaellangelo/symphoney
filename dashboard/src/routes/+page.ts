import type { PageLoad } from "./projects/$types";

export const load: PageLoad = () => {
    return {
        projectCount: Math.floor(Math.random() * 100)
    }
}