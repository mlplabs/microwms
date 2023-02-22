import { createRouter, createWebHashHistory } from 'vue-router' //
import MainView from "@/view/MainView";
import ReferencesView from "@/view/ReferencesView";
import ReferenceProducts from "@/view/refs/ReferenceProducts";
import ReferenceManufacturers from "@/view/refs/ReferenceManufacturers";
import ReferenceWhs from "@/view/refs/ReferenceWhs";
import ReferenceBarcodes from "@/view/refs/ReferenceBarcodes";
import ReportsView from "@/view/ReportsView";
import PropertiesView from "@/view/PropertiesView";
import GlobalSearchView from "@/view/GlobalSearchView";
import DocShipment from "@/view/docs/DocShipment";
import DocReceipt from "@/view/docs/DocReceipt";
import HardwarePrinters from "@/view/hardware/preference/HardwarePrinters";
import ReferenceUsers from "@/view/refs/ReferenceUsers";

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
        name: 'ReferenceProducts',
        component: ReferenceProducts
    },
    {
        path: '/refs/manufacturers',
        name: 'ReferenceManufacturers',
        component: ReferenceManufacturers
    },
    {
        path: '/refs/whs',
        name: 'ReferenceWhs',
        component: ReferenceWhs
    },
    {
        path: '/refs/barcodes',
        name: 'ReferenceBarcodes',
        component: ReferenceBarcodes
    },
    {
        path: '/refs/users',
        name: 'ReferenceUsers',
        component: ReferenceUsers
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

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router