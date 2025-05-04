import Link from "next/link"
import { FileCheck, Plus, Search } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Header } from "@/components/header"
import { Footer } from "@/components/footer"

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      <main className="flex-1">
        <section className="bg-muted py-12 md:py-24 lg:py-32">
          <div className="container px-4 md:px-6">
            <div className="flex flex-col items-center justify-center space-y-4 text-center">
              <div className="space-y-2">
                <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl">
                  Secure Certificate Management on Blockchain
                </h1>
                <p className="mx-auto max-w-[700px] text-muted-foreground md:text-xl">
                  Issue, verify, and manage certificates with immutable blockchain technology
                </p>
              </div>
              <div className="flex flex-col gap-2 min-[400px]:flex-row">
                <Link href="/issue">
                  <Button>
                    <Plus className="mr-2 h-4 w-4" />
                    Issue Certificate
                  </Button>
                </Link>
                <Link href="/verify">
                  <Button variant="outline">
                    <FileCheck className="mr-2 h-4 w-4" />
                    Verify Certificate
                  </Button>
                </Link>
              </div>
            </div>
          </div>
        </section>
        <section className="container py-12 md:py-24 lg:py-32">
          <div className="mx-auto grid max-w-5xl items-center gap-6 py-12 lg:grid-cols-2 lg:gap-12">
            <div className="space-y-4">
              <h2 className="text-3xl font-bold tracking-tighter md:text-4xl">How It Works</h2>
              <p className="text-muted-foreground md:text-xl">
                Our platform uses blockchain technology to create tamper-proof digital certificates that can be easily
                verified.
              </p>
              <ul className="grid gap-2">
                <li className="flex items-center gap-2">
                  <div className="flex h-8 w-8 items-center justify-center rounded-full bg-primary text-primary-foreground">
                    1
                  </div>
                  <span>Issue certificates with unique identifiers</span>
                </li>
                <li className="flex items-center gap-2">
                  <div className="flex h-8 w-8 items-center justify-center rounded-full bg-primary text-primary-foreground">
                    2
                  </div>
                  <span>Store certificate data securely on the blockchain</span>
                </li>
                <li className="flex items-center gap-2">
                  <div className="flex h-8 w-8 items-center justify-center rounded-full bg-primary text-primary-foreground">
                    3
                  </div>
                  <span>Verify certificates instantly with our verification tool</span>
                </li>
              </ul>
            </div>
            <div className="rounded-lg border bg-background p-8">
              <Tabs defaultValue="issue">
                <TabsList className="grid w-full grid-cols-2">
                  <TabsTrigger value="issue">Issue</TabsTrigger>
                  <TabsTrigger value="verify">Verify</TabsTrigger>
                </TabsList>
                <TabsContent value="issue" className="space-y-4 pt-4">
                  <div className="space-y-2">
                    <h3 className="text-lg font-medium">Issue a New Certificate</h3>
                    <p className="text-sm text-muted-foreground">
                      Create and issue a new certificate to be stored on the blockchain.
                    </p>
                  </div>
                  <Link href="/issue">
                    <Button className="w-full">
                      <Plus className="mr-2 h-4 w-4" />
                      Issue Certificate
                    </Button>
                  </Link>
                </TabsContent>
                <TabsContent value="verify" className="space-y-4 pt-4">
                  <div className="space-y-2">
                    <h3 className="text-lg font-medium">Verify a Certificate</h3>
                    <p className="text-sm text-muted-foreground">
                      Enter a certificate ID or hash to verify its authenticity.
                    </p>
                  </div>
                  <div className="flex space-x-2">
                    <Input placeholder="Enter certificate ID" />
                    <Button>
                      <Search className="h-4 w-4" />
                    </Button>
                  </div>
                </TabsContent>
              </Tabs>
            </div>
          </div>
        </section>
        <section className="bg-muted py-12 md:py-24 lg:py-32">
          <div className="container px-4 md:px-6">
            <div className="mx-auto grid max-w-5xl gap-6 md:grid-cols-2 lg:grid-cols-3">
              <Card>
                <CardHeader>
                  <CardTitle>For Educational Institutions</CardTitle>
                  <CardDescription>Issue verifiable certificates for your students</CardDescription>
                </CardHeader>
                <CardContent>
                  <p>
                    Provide your graduates with tamper-proof digital certificates that can be easily shared and verified
                    by employers.
                  </p>
                </CardContent>
                <CardFooter>
                  <Button variant="outline" className="w-full">
                    Learn More
                  </Button>
                </CardFooter>
              </Card>
              <Card>
                <CardHeader>
                  <CardTitle>For Employers</CardTitle>
                  <CardDescription>Verify candidate credentials instantly</CardDescription>
                </CardHeader>
                <CardContent>
                  <p>
                    Quickly verify the authenticity of certificates presented by job applicants with our simple
                    verification tool.
                  </p>
                </CardContent>
                <CardFooter>
                  <Button variant="outline" className="w-full">
                    Learn More
                  </Button>
                </CardFooter>
              </Card>
              <Card>
                <CardHeader>
                  <CardTitle>For Certificate Holders</CardTitle>
                  <CardDescription>Manage and share your credentials</CardDescription>
                </CardHeader>
                <CardContent>
                  <p>
                    Access all your certificates in one place and share them with potential employers with confidence.
                  </p>
                </CardContent>
                <CardFooter>
                  <Button variant="outline" className="w-full">
                    Learn More
                  </Button>
                </CardFooter>
              </Card>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  )
}