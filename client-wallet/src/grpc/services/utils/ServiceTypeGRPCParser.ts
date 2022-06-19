import { ServiceTypeEnum } from '../../pb/client_wallet_pb'
import { ServiceType } from '../../../core/transaction/Transaction'

export default class ServiceTypeGRPCParser {
  public static toDomain (service: ServiceTypeEnum): ServiceType {
    switch (service) {
      case ServiceTypeEnum.CARD_PAYMENT:
        return ServiceType.CARD_PAYMENT
      case ServiceTypeEnum.INTERNAL:
        return ServiceType.INTERNAL
      default:
        throw new Error(`Unknown service: ${service}`)
    }
  }

  public static toGRPC (service: ServiceType): ServiceTypeEnum {
    switch (service) {
      case ServiceType.CARD_PAYMENT:
        return ServiceTypeEnum.CARD_PAYMENT
      case ServiceType.INTERNAL:
        return ServiceTypeEnum.INTERNAL
      default:
        throw new Error(`Unknown service: ${service}`)
    }
  }
}
