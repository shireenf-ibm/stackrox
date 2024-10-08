import React from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';

import usePageAction from 'hooks/usePageAction';
import usePermissions from 'hooks/usePermissions';
import { complianceEnhancedSchedulesPath } from 'routePaths';

import { scanConfigDetailsPath } from './compliance.scanConfigs.routes';
import { PageActions } from './compliance.scanConfigs.utils';
import CreateScanConfigPage from './CreateScanConfigPage';
import ScanConfigDetailPage from './ScanConfigDetailPage';
import ScanConfigsTablePage from './ScanConfigsTablePage';

function ScanConfigsPage() {
    /*
     * Examples of urls for ScanConfigPage:
     * /main/compliance-enhanced/scan-configs
     * /main/compliance-enhanced/scan-configs?action=create
     * /main/compliance-enhanced/scan-configs/configId
     */
    const { pageAction } = usePageAction<PageActions>();

    const { hasReadWriteAccess } = usePermissions();
    const hasWriteAccessForCompliance = hasReadWriteAccess('Compliance');

    return (
        <Switch>
            <Route
                exact
                path={complianceEnhancedSchedulesPath}
                render={() => {
                    if (pageAction === 'create' && hasWriteAccessForCompliance) {
                        return <CreateScanConfigPage />;
                    }
                    if (!pageAction) {
                        return (
                            <ScanConfigsTablePage
                                hasWriteAccessForCompliance={hasWriteAccessForCompliance}
                            />
                        );
                    }
                    return <Redirect to={complianceEnhancedSchedulesPath} />;
                }}
            />
            <Route exact path={scanConfigDetailsPath}>
                <ScanConfigDetailPage hasWriteAccessForCompliance={hasWriteAccessForCompliance} />
            </Route>
        </Switch>
    );
}

export default ScanConfigsPage;
