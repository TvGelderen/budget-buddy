export type TransactionForm = {
    amount: number;
    description: string;
    startDate: Date;
    endDate: Date | null;
    recurring: boolean;
    interval: string | null;
    daysInterval: number | null;
    incoming: boolean;
    type: string | null;
    errors: TransactionFormErrors;
};

export type TransactionFormErrors = {
    amount: string | null;
    description: string | null;
    startDate: string | null;
    endDate: string | null;
    interval: string | null;
    daysInterval: string | null;
    type: string | null;
};