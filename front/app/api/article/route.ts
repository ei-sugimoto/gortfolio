import { NextResponse } from "next/server";
import config from "@/config";

export type ArticleResponse = {
  title: string;
  url: string;
};

export async function GET() {
  try {
    const url = `${config.url}/article`;
    const res = await fetch(url);
    const data = await res.json();
    return NextResponse.json(data as ArticleResponse[]);
  } catch (error) {
    return NextResponse.error();
  }
}
