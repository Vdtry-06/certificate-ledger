const API_BASE_URL = "http://localhost:8080/api"

import { getToken } from "./auth"

async function fetchAPI<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const url = `${API_BASE_URL}${endpoint}`

  const token = getToken()

  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers,
  }

  try {
    const response = await fetch(url, {
      ...options,
      headers,
    })

    if (!response.ok) {
      try {
        const errorData = await response.json()
        throw new Error(errorData.message || `API error: ${response.status}`)
      } catch {
        throw new Error(`API error: ${response.status}`)
      }
    }

    if (response.status === 204) {
      return {} as T
    }

    return await response.json()
  } catch (error) {
    console.error("API request failed:", error)
    throw error
  }
}

export interface Certificate {
  id: string
  hash: string
  recipientName: string
  recipientEmail: string
  certificateTitle: string
  issueDate: string
  issuerName: string
  description: string
  blockNumber: number
  timestamp: string
  issuerId?: string
}

export interface CertificateRequest {
  recipientName: string
  recipientEmail: string
  certificateTitle: string
  issueDate: string
  issuerName: string
  description: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  name: string
  email: string
  password: string
}

export interface AuthResponse {
  token: string
  user: {
    id: string
    name: string
    email: string
    createdAt: string
    updatedAt: string
  }
}

export async function createCertificate(data: CertificateRequest): Promise<Certificate> {
  return fetchAPI<Certificate>("/certificates", {
    method: "POST",
    body: JSON.stringify(data),
  })
}

export async function getCertificate(id: string): Promise<Certificate> {
  return fetchAPI<Certificate>(`/certificates/${id}`)
}

export async function verifyCertificate(hash: string): Promise<{ valid: boolean }> {
  return fetchAPI<{ valid: boolean }>(`/certificates/verify/${hash}`)
}

export async function getAllCertificates(): Promise<Certificate[]> {
  return fetchAPI<Certificate[]>("/certificates")
}

export async function login(credentials: LoginRequest): Promise<AuthResponse> {
  return fetchAPI<AuthResponse>("/auth/login", {
    method: "POST",
    body: JSON.stringify(credentials),
  })
}

export async function register(userData: RegisterRequest): Promise<AuthResponse> {
  return fetchAPI<AuthResponse>("/auth/register", {
    method: "POST",
    body: JSON.stringify(userData),
  })
}

export async function getUserCertificates(userId: string): Promise<Certificate[]> {
  return fetchAPI<Certificate[]>(`/users/${userId}/certificates`)
}

export async function getUserIssuedCertificates(): Promise<Certificate[]> {
  return fetchAPI<Certificate[]>("/certificates")
}

export async function getUserReceivedCertificates(email: string): Promise<Certificate[]> {
  const allCerts = await getAllCertificates()
  return allCerts.filter((cert) => cert.recipientEmail === email)
}
