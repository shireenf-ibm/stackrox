import React, { ReactElement } from 'react';
import { Alert, Flex, FlexItem } from '@patternfly/react-core';

import Loader from 'Components/Loader';
import { labelClassName } from 'constants/form.constants';
import { Cluster, ClusterManagerType } from 'types/cluster.proto';
import { DecommissionedClusterRetentionInfo } from 'types/clusterService.proto';

import ClusterSummary from './Components/ClusterSummary';
import StaticConfigurationSection from './StaticConfigurationSection';
import DynamicConfigurationSection from './DynamicConfigurationSection';
import ClusterLabelsTable from './ClusterLabelsTable';

type ClusterEditFormProps = {
    centralVersion: string;
    clusterRetentionInfo: DecommissionedClusterRetentionInfo;
    selectedCluster: Cluster;
    managerType: ClusterManagerType;
    handleChange: (any) => void;
    handleChangeLabels: (labels) => void;
    isLoading: boolean;
};

function ClusterEditForm({
    centralVersion,
    clusterRetentionInfo,
    selectedCluster,
    managerType,
    handleChange,
    handleChangeLabels,
    isLoading,
}: ClusterEditFormProps): ReactElement {
    if (isLoading) {
        return <Loader />;
    }
    const isManagerTypeNonConfigurable =
        managerType === 'MANAGER_TYPE_KUBERNETES_OPERATOR' ||
        managerType === 'MANAGER_TYPE_HELM_CHART';
    return (
        <div className="bg-base-200 px-4 w-full">
            {/* @TODO, replace open prop with dynamic logic, based on clusterType */}
            {selectedCluster.id && selectedCluster.healthStatus ? (
                <ClusterSummary
                    healthStatus={selectedCluster.healthStatus}
                    status={selectedCluster.status}
                    centralVersion={centralVersion}
                    clusterId={selectedCluster.id}
                    autoRefreshEnabled={selectedCluster.sensorCapabilities?.includes(
                        'SecuredClusterCertificatesRefresh'
                    )}
                    clusterRetentionInfo={clusterRetentionInfo}
                    isManagerTypeNonConfigurable={isManagerTypeNonConfigurable}
                />
            ) : (
                <Alert variant="warning" isInline title="Legacy installation method" component="p">
                    <Flex direction={{ default: 'column' }}>
                        <FlexItem>
                            <p>
                                To avoid extra operational complexity, use a{' '}
                                <strong>cluster init bundle</strong> with either of the following
                                installation methods:
                            </p>
                            <p>
                                <strong>Operator</strong> for Red Hat OpenShift
                            </p>
                            <p>
                                <strong>Helm chart</strong> for other platforms
                            </p>
                        </FlexItem>
                        <FlexItem>
                            <p>
                                Only use the legacy installation method if you have a specific
                                installation need that requires using this method.
                            </p>
                        </FlexItem>
                    </Flex>
                </Alert>
            )}
            <form
                className="grid grid-columns-1 md:grid-columns-2 grid-gap-4 xl:grid-gap-6 mb-4 w-full"
                data-testid="cluster-form"
            >
                <StaticConfigurationSection
                    isManagerTypeNonConfigurable={isManagerTypeNonConfigurable}
                    handleChange={handleChange}
                    selectedCluster={selectedCluster}
                />
                <div>
                    <DynamicConfigurationSection
                        dynamicConfig={selectedCluster.dynamicConfig}
                        helmConfig={selectedCluster.helmConfig}
                        handleChange={handleChange}
                        clusterType={selectedCluster.type}
                        isManagerTypeNonConfigurable={isManagerTypeNonConfigurable}
                    />
                    <div className="pt-4">
                        <label htmlFor="labels" className={labelClassName}>
                            Cluster labels
                        </label>
                        <ClusterLabelsTable
                            labels={selectedCluster?.labels ?? {}}
                            handleChangeLabels={handleChangeLabels}
                            hasAction
                        />
                    </div>
                </div>
            </form>
        </div>
    );
}

export default ClusterEditForm;
