"use client"

import type React from "react"

import { useState } from "react"
import Link from "next/link"
import { Award, ChevronLeft } from "lucide-react"
import { toast } from "sonner"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { createCertificate, type CertificateRequest } from "@/lib/api"
import { Header } from "@/components/header"
import { Footer } from "@/components/footer"

export default function IssueCertificate() {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [formData, setFormData] = useState<CertificateRequest>({
    recipientName: "",
    recipientEmail: "",
    certificateTitle: "",
    issueDate: "",
    description: "",
    issuerName: "",
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setFormData((prev) => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)

    try {
      const certificate = await createCertificate(formData)

      toast.success("Certificate Issued", {
        description: `Certificate for ${certificate.recipientName} has been successfully issued and added to the blockchain.`,
      })

      setFormData({
        recipientName: "",
        recipientEmail: "",
        certificateTitle: "",
        issueDate: "",
        description: "",
        issuerName: "",
      })
    } catch (error) {
      toast.error("Error", {
        description: error instanceof Error ? error.message : "Failed to issue certificate. Please try again.",
      })
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      <div className="container max-w-4xl px-4 mx-auto py-10">
        <div className="mb-8">
          <Link href="/" className="flex items-center text-sm text-muted-foreground hover:text-foreground">
            <ChevronLeft className="mr-1 h-4 w-4" />
            Back to Home
          </Link>
        </div>
        <div className="flex items-center gap-2 mb-8">
          <Award className="h-8 w-8" />
          <h1 className="text-3xl font-bold">Issue Certificate</h1>
        </div>
        <Card>
          <CardHeader>
            <CardTitle>Certificate Details</CardTitle>
            <CardDescription>
              Fill in the details of the certificate you want to issue. Once issued, the certificate will be permanently
              stored on the blockchain.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="grid gap-4 md:grid-cols-2">
                <div className="space-y-2">
                  <Label htmlFor="recipientName">Recipient Name</Label>
                  <Input
                    id="recipientName"
                    name="recipientName"
                    value={formData.recipientName}
                    onChange={handleChange}
                    placeholder="John Doe"
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="recipientEmail">Recipient Email</Label>
                  <Input
                    id="recipientEmail"
                    name="recipientEmail"
                    type="email"
                    value={formData.recipientEmail}
                    onChange={handleChange}
                    placeholder="john.doe@example.com"
                    required
                  />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="certificateTitle">Certificate Title</Label>
                <Input
                  id="certificateTitle"
                  name="certificateTitle"
                  value={formData.certificateTitle}
                  onChange={handleChange}
                  placeholder="e.g., Blockchain Development Certification"
                  required
                />
              </div>
              <div className="grid gap-4 md:grid-cols-2">
                <div className="space-y-2">
                  <Label htmlFor="issueDate">Issue Date</Label>
                  <Input
                    id="issueDate"
                    name="issueDate"
                    type="date"
                    value={formData.issueDate}
                    onChange={handleChange}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="issuerName">Issuer Name</Label>
                  <Input
                    id="issuerName"
                    name="issuerName"
                    value={formData.issuerName}
                    onChange={handleChange}
                    placeholder="Your Organization"
                    required
                  />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="description">Description</Label>
                <Textarea
                  id="description"
                  name="description"
                  value={formData.description}
                  onChange={handleChange}
                  placeholder="Describe the achievement or qualification this certificate represents"
                  rows={4}
                  required
                />
              </div>
              <Button type="submit" className="w-full" disabled={isSubmitting}>
                {isSubmitting ? "Issuing Certificate..." : "Issue Certificate"}
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>
      <Footer />
    </div>
  )
}
