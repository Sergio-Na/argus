import { useState } from 'react'
import { json, ActionFunctionArgs } from '@remix-run/node'
import { Form, useActionData, Link } from '@remix-run/react'
import { Button } from "../components/ui/button"
import { Input } from "../components/ui/input"
import { Label } from "../components/ui/label"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "../components/ui/card"
import { EyeIcon, EyeOffIcon } from 'lucide-react'
import { login } from '../lib/api/helpers'

export const action = async ({ request }: ActionFunctionArgs) => {
    const formData = await request.formData()
    const email = formData.get('email') as string
    const password = formData.get('password') as string
    const response = await login(email, password)
    console.log(response)
    // Here you would typically validate the credentials and log the user in
    // For this example, we'll just return a success message
    return json({ success: true, message: 'Login successful' })
}

export default function Login() {
    const [showPassword, setShowPassword] = useState(false)
    const actionData = useActionData<typeof action>()

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <Card className="w-full max-w-md">
                <CardHeader>
                    <CardTitle className="text-2xl font-bold text-center">Login</CardTitle>
                    <CardDescription className="text-center">Welcome back! Please login to your account</CardDescription>
                </CardHeader>
                <Form method="post">
                    <CardContent>
                        <div className="space-y-4">
                            <div className="space-y-2">
                                <Label htmlFor="email">Email</Label>
                                <Input id="email" name="email" type="email" placeholder="Enter your email" required />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="password">Password</Label>
                                <div className="relative">
                                    <Input
                                        id="password"
                                        name="password"
                                        type={showPassword ? "text" : "password"}
                                        placeholder="Enter your password"
                                        required
                                    />
                                    <Button
                                        type="button"
                                        variant="ghost"
                                        size="sm"
                                        className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                                        onClick={() => setShowPassword(!showPassword)}
                                        aria-label={showPassword ? "Hide password" : "Show password"}
                                    >
                                        {showPassword ? (
                                            <EyeOffIcon className="h-4 w-4 text-gray-500" />
                                        ) : (
                                            <EyeIcon className="h-4 w-4 text-gray-500" />
                                        )}
                                    </Button>
                                </div>
                            </div>
                        </div>
                    </CardContent>
                    <CardFooter className="flex flex-col space-y-4">
                        <Button className="w-full" type="submit">Login</Button>
                        <div className="text-sm text-center">
                            No account?
                            <Link to="/signup" className="text-primary hover:underline ml-1">
                                Sign up
                            </Link>
                        </div>
                    </CardFooter>
                </Form>
                {actionData?.success && (
                    <p className="text-center text-green-600 mt-4">{actionData.message}</p>
                )}
            </Card>
        </div>
    )
}