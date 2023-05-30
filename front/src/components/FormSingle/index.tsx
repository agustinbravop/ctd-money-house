import { Button, FormControl, InputLabel, OutlinedInput } from '@mui/material';
import React, { useMemo } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { CardCustom, ErrorMessage, Errors } from '..';
import { valuesHaveErrors } from '../../utils';

interface SendMoneyInputProps {
    formState: any;
    label: string;
    name: string;
    title: string;
    actionLabel: string;
    validation?: any;
    submit: () => void;
    handleChange: (
        event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => void;
    type: string;
}

export const FormSingle = ({
                               formState,
                               label,
                               name,
                               title,
                               type,
                               actionLabel,
                               handleChange,
                               validation,
                               submit,
                           }: SendMoneyInputProps) => {
    const {
        register,
        handleSubmit,
        formState: { errors, isDirty },
    } = useForm({
        criteriaMode: 'all',
    });

    const isEmpty = formState[name] === '';
    const hasErrors = useMemo(() => valuesHaveErrors(errors), [errors]);

    const onSubmit: SubmitHandler<any> = () => {
        submit();
    };

    return (
        <CardCustom
            className="tw-max-w-5xl"
            content={
                <>
                    <div>
                        <div className="tw-flex tw-justify-between tw-mb-8">
                            <p className="tw-font-bold">{title}</p>
                        </div>
                        <form onSubmit={handleSubmit(onSubmit)}>
                            <FormControl variant="outlined">
                                <InputLabel htmlFor={`outlined-adornment-${name}`}>
                                    {label}
                                </InputLabel>
                                <OutlinedInput
                                    id={`outlined-adornment-${name}`}
                                    type={type}
                                    value={formState[name]}
                                    {...register(name, validation)}
                                    onChange={handleChange}
                                    label={name}
                                    autoComplete="off"
                                />
                            </FormControl>
                            {errors[name] && <ErrorMessage errors={errors[name] as Errors} />}
                            <div className="tw-flex tw-w-full tw-justify-end tw-mt-6">
                                <Button
                                    className={`tw-h-12 tw-w-64 ${
                                        hasErrors || isEmpty || !isDirty
                                            ? 'tw-text-neutral-gray-300 tw-border-neutral-gray-300 tw-cursor-not-allowed'
                                            : ''
                                    }`}
                                    variant="outlined"
                                    disabled={hasErrors || isEmpty || !isDirty}
                                    type="submit"
                                >
                                    {actionLabel}
                                </Button>
                            </div>
                        </form>
                    </div>
                </>
            }
        />
    );
};
