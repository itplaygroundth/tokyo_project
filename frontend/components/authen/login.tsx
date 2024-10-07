'use client'
import { Input } from "@/components/ui/input"
import useAuthStore from "@/store/auth" 
import { useForm, SubmitHandler } from "react-hook-form"

type User = {
  email:string
  password:string
}

export default function Component() {

  const {
    register,
    handleSubmit
  } = useForm<User>()

  const {login,logout} = useAuthStore()
  const onSubmit: SubmitHandler<User> = (data:User) => {
  
  console.log('logged in')
  login()
 
  }

  // const handleLogOut() => {
  //     console.log('logged out')
  //     logout()
  //     }
  return (
    <>
    <form onSubmit={handleSubmit(onSubmit)}>
    <div className="bg-gray-100 min-h-screen flex items-center justify-center p-6">
      <div className="bg-white shadow-lg rounded-lg max-w-md mx-auto">
        <div className="px-6 py-4">
          <h2 className="text-gray-700 text-3xl font-semibold">Login</h2>
          <p className="mt-1 text-gray-600">Please login to your account.</p>
        </div>
        <div className="px-6 py-4">
            <div className="mt-4">
              <label className="block text-gray-700" htmlFor="email">
                Email
              </label>
              <Input
                type="email"
                id="email"
                className="mt-2 rounded w-full px-3 py-2 text-gray-700 bg-gray-200 outline-none focus:bg-gray-300"
                placeholder="m@example.com"
                required
                defaultValue="m@example.com"  {...register("email", { required: true })} 
              />
            </div>
            <div className="mt-4">
              <label className="block text-gray-700" htmlFor="password">
                Password
              </label>
              <Input
                type="password"
                id="password"
                className="mt-2 rounded w-full px-3 py-2 text-gray-700 bg-gray-200 outline-none focus:bg-gray-300"
                required
                defaultValue="" {...register("password", { required: true })} 
              />
            </div>
            <div className="mt-6">
              <button type="submit" className="py-2 px-4 bg-gray-700 text-white rounded hover:bg-gray-600 w-full">
                Login
              </button>
            </div>
        </div>
      </div>
    </div>
    </form>
    </>
  )
}