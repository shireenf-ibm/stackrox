import React, { ReactElement } from 'react';
import {
    Alert,
    Form,
    PageSection,
    Text,
    TextInput,
    ToggleGroup,
    ToggleGroupItem,
} from '@patternfly/react-core';
import * as yup from 'yup';
import merge from 'lodash/merge';

import { ClairifyImageIntegration } from 'types/imageIntegration.proto';

import FormMessage from 'Components/PatternFly/FormMessage';
import FormTestButton from 'Components/PatternFly/FormTestButton';
import FormSaveButton from 'Components/PatternFly/FormSaveButton';
import FormCancelButton from 'Components/PatternFly/FormCancelButton';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormLabelGroup from '../FormLabelGroup';

import { categoriesUtilsForClairifyScanner } from '../../utils/integrationUtils';

const { categoriesAlternatives, getCategoriesText, matchCategoriesAlternative, validCategories } =
    categoriesUtilsForClairifyScanner;

export const validationSchema = yup.object().shape({
    name: yup.string().trim().required('An integration name is required'),
    categories: yup
        .array()
        .of(yup.string().trim().oneOf(validCategories))
        .min(1, 'Must have at least one type selected')
        .required('A category is required'),
    clairify: yup.object().shape({
        endpoint: yup.string().trim().required('An endpoint is required'),
        grpcEndpoint: yup.string().trim(),
        numConcurrentScans: yup.string().trim(),
    }),
    type: yup.string().matches(/clairify/),
});

export const defaultValues: ClairifyImageIntegration = {
    id: '',
    name: '',
    categories: ['SCANNER'],
    clairify: {
        endpoint: '',
        grpcEndpoint: '',
        numConcurrentScans: 0,
    },
    autogenerated: false,
    clusterId: '',
    skipTestIntegration: false,
    type: 'clairify',
};

function ClairifyIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<ClairifyImageIntegration>): ReactElement {
    const formInitialValues: ClairifyImageIntegration = merge({}, defaultValues, initialValues);
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<ClairifyImageIntegration>({
        initialValues: formInitialValues,
        validationSchema,
    });

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                <Alert
                    title="Deprecation notice"
                    component="p"
                    variant={'warning'}
                    isInline
                    className="pf-v5-u-mb-lg"
                >
                    <Text>StackRox Scanner will be removed in a future release.</Text>
                    <Text>No new enhancements for StackRox Scanner will be done or accepted.</Text>
                    <Text>It is recommended to use Scanner V4, instead.</Text>
                </Alert>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="name"
                            value={values.name}
                            placeholder="(example, StackRox Scanner Integration)"
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Type"
                        isRequired
                        fieldId="categories"
                        touched={touched}
                        errors={errors}
                    >
                        <ToggleGroup id="categories" areAllGroupsDisabled={!isEditable}>
                            {categoriesAlternatives.map((categoriesAlternative) => {
                                const [categoriesAlternativeItem0] = categoriesAlternative;
                                const text = getCategoriesText(categoriesAlternativeItem0);
                                const isSelected = matchCategoriesAlternative(
                                    categoriesAlternative,
                                    values.categories
                                );
                                return (
                                    <ToggleGroupItem
                                        key={text}
                                        text={text}
                                        isSelected={isSelected}
                                        onChange={() =>
                                            setFieldValue('categories', categoriesAlternativeItem0)
                                        }
                                    />
                                );
                            })}
                        </ToggleGroup>
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Endpoint"
                        isRequired
                        fieldId="clairify.endpoint"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="clairify.endpoint"
                            value={values.clairify.endpoint}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="GRPC endpoint"
                        fieldId="clairify.grpcEndpoint"
                        helperText="Used for node scanning"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="clairify.grpcEndpoint"
                            value={values.clairify.grpcEndpoint}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Max concurrent image scans"
                        fieldId="clairify.numConcurrentScans"
                        helperText="0 for default"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="number"
                            id="clairify.numConcurrentScans"
                            value={values.clairify.numConcurrentScans}
                            onChange={(event, value) => onChange(value, event)}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default ClairifyIntegrationForm;
