import { createRouter, createWebHashHistory } from 'vue-router' //
import MainView from "@/view/MainView";
import ReferencesView from "@/view/ReferencesView";
import CatalogProducts from "@/view/refs/CatalogProducts";
import CatalogManufacturers from "@/view/refs/CatalogManufacturers";
import ReferenceWhs from "@/view/refs/CatalogWhs";
import CatalogBarcodes from "@/view/refs/CatalogBarcodes";
import ReportsView from "@/view/ReportsView";
import PropertiesView from "@/view/PropertiesView";
import GlobalSearchView from "@/view/GlobalSearchView";
import DocShipment from "@/view/docs/DocShipment";
import DocReceipt from "@/view/docs/DocReceipt";
import HardwarePrinters from "@/view/hardware/preference/HardwarePrinters";
import CatalogUsers from "@/view/refs/CatalogUsers";
import RemainingProducts from "@/view/reports/RemainingProducts"
import HistoryProducts from "@/view/reports/HistoryProducts"

const routes = [
    {
        path: '/',
        name: 'MainView',
        component: MainView
    },
    {
        path: '/hw/pref',
        name: 'HardwarePrinters',
        component: HardwarePrinters
    },
    {
        path: '/refs',
        name: 'ReferencesView',
        component: ReferencesView
    },
    {
        path: '/refs/products',
        name: 'CatalogProducts',
        component: CatalogProducts
    },
    {
        path: '/refs/manufacturers',
        name: 'CatalogManufacturers',
        component: CatalogManufacturers
    },
    {
        path: '/refs/whs',
        name: 'ReferenceWhs',
        component: ReferenceWhs
    },
    {
        path: '/refs/barcodes',
        name: 'CatalogBarcodes',
        component: CatalogBarcodes
    },
    {
        path: '/refs/users',
        name: 'CatalogUsers',
        component: CatalogUsers
    },
    {
        path: '/docs/shipment',
        name: 'DocShipment',
        component: DocShipment
    },
    {
        path: '/docs/receipt',
        name: 'DocReceipt',
        component: DocReceipt
    },
    {
        path: '/reports',
        name: 'ReportsView',
        component: ReportsView
    },
    {
        path: '/reports/remaining/:propProductId',
        name: 'remaining',
        component: RemainingProducts,
        props: castRouteParams
    },
    {
        path: '/reports/history/:propProductId',
        name: 'history',
        component: HistoryProducts,
        props: castRouteParams
    },

    {
        path: '/props',
        name: 'PropertiesView',
        component: PropertiesView
    },
    {
        path: '/search/:searchData',
        name: 'GlobalSearchView',
        component: GlobalSearchView,
        props: true

    },
]

function castRouteParams(route) {
    return {
        propProductId: Number(route.params.propProductId),
    };
}

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router