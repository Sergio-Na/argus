import type { EmailSignUpData, EmailLoginData } from "./types";
const API_BASEURL = process.env.API_BASEURL;

export async function signUp(email: string, password: string): Promise<JSON> {
  const signUpData: EmailSignUpData = {
    email,
    password,
  };

  try {
    const response = await fetch(`${API_BASEURL}/signUp`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(signUpData),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Signup error:", error);
    throw error;
  }
}

export async function login(email: string, password: string): Promise<JSON> {
  const loginData: EmailLoginData = {
    email,
    password,
  };

  try {
    const response = await fetch(`${API_BASEURL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(loginData),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Signup error:", error);
    throw error;
  }
}
