import { type RequestHandler } from "@sveltejs/kit";
import { authorizeFetch, authorizeFetchBody, forbidden, toISOString } from "$lib";
import type { TransactionForm } from "../../../../ambient";
import { verifyForm } from "$lib/api/transactions";

export const PUT: RequestHandler = async ({ locals: { session }, request, params }) => {
    if (!session) {
        return forbidden();
    }

    const id = params["id"];
    if (id === null) {
        return new Response(null, {
            status: 400,
        })
    }

    const form: TransactionForm = await request.json();
    const errors = verifyForm(form);
    const isValid = Object.values(errors).every((err) => err === null);
    if (!isValid) {
        form.errors = errors;
        return new Response(JSON.stringify(form), {
            status: 400,
            headers: { "Content-Type": "application/json" },
        });
    }

    form.startDate = toISOString(form.startDate!);
    form.endDate = form.recurring ? toISOString(form.endDate!) : form.startDate;

    const response = await authorizeFetchBody(
        `transactions/${id}`,
        session?.accessToken,
        "PUT",
        JSON.stringify(form));
    if (response.ok) {
        return response;
    }

    return new Response(null, {
        status: 500,
    })
}

export const DELETE: RequestHandler = async ({ locals: { session }, params }) => {
    if (!session) {
        return forbidden();
    }

    const id = params["id"];
    if (id === null) {
        return new Response(null, {
            status: 400,
        })
    }

    const response = await authorizeFetch(`transactions/${id}`, session?.accessToken, "DELETE");
    if (response.ok) {
        return response;
    }

    return new Response(null, {
        status: 500,
    })
}

