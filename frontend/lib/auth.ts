export function isLoggedIn(): boolean {
    if (typeof window === "undefined") return false
    return !!localStorage.getItem("token")
  }
  
  export function getCurrentUser() {
    if (typeof window === "undefined") return null
  
    const userJson = localStorage.getItem("user")
    if (!userJson) return null
  
    try {
      return JSON.parse(userJson)
    } catch {
      return null
    }
  }
  
  export function getToken(): string | null {
    if (typeof window === "undefined") return null
    return localStorage.getItem("token")
  }
  
  export function setAuthData(token: string, user: any): void {
    localStorage.setItem("token", token)
    localStorage.setItem("user", JSON.stringify(user))
  }
  
  export function logout(): void {
    if (typeof window === "undefined") return
    localStorage.removeItem("token")
    localStorage.removeItem("user")
    window.location.href = "/login"
  }
  