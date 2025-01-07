import { NextRequest, NextResponse } from "next/server";

// Define the base URL of your Golang backend
const GO_BACKEND_URL = process.env.GO_BACKEND_URL || "http://localhost:8080";

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

 // Check if the request should be proxied to the Golang backend
  if (pathname.startsWith("/api/")) {
    
    const proxyUrl = `${GO_BACKEND_URL}${pathname}`;
    // Forward the request to the Golang backend
    try {
      console.log(req.method)
      const response = await fetch(proxyUrl,
        {method: req.method,body: req.method !== "GET" ?  req.body : undefined,});
     
      
      // Return the response received from the Golang backend
      return new NextResponse(  response.body, {
        status:   response.status,
        headers:   req.headers})


    } catch (error) {
      console.log("error")
      console.error(error);
    }
  
  }

  // Continue with Next.js routing for non-API requests
  return NextResponse.next();
}
