import React, { ReactElement } from 'react';
import { Alert } from '@patternfly/react-core';

export type FormResponseMessage = {
    message: string;
    isError: boolean;
} | null;

export type FormMessageProps = {
    message: FormResponseMessage;
};

function FormMessage({ message }: FormMessageProps): ReactElement {
    const title = message?.isError ? 'Failure' : 'Success';
    const variant = message?.isError ? 'danger' : 'success';
    return (
        <div id="form-message-alert">
            {message && (
                <Alert
                    className="pf-v5-u-mt-md pf-v5-u-mb-md"
                    title={title}
                    component="p"
                    variant={variant}
                    isInline
                >
                    <p>{message?.message}</p>
                </Alert>
            )}
        </div>
    );
}

export default FormMessage;
