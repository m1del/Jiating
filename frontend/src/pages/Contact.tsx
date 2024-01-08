import { EmailForm } from '../components'

function Contact() {
  return (
    <div className='min-h-screen container mx-auto p-4'>

      <div className="bg-gray-700 text-white p-6 rounded-md shadow-md min-w-full">
          <h1 className="text-4xl font-bold mb-2">Contact Us</h1>
          <p className="text-xl">Be a part of the community or schedule a performance!</p>
      </div>

      <div className='p-4'>
        <EmailForm/>
      </div>

    </div>
  )
}

export default Contact
