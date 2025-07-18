"use client"

import { Check, X, Star, Zap, Shield, Users } from "lucide-react"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"

const plans = [
  {
    name: "Básico",
    price: "R$ 29",
    period: "/mês",
    description: "Ideal para iniciantes",
    popular: false,
    features: {
      users: "1 usuário",
      storage: "10 GB",
      support: "Email",
      analytics: false,
      api: false,
      customDomain: false,
      ssl: true,
      backup: "Semanal",
    },
  },
  {
    name: "Profissional",
    price: "R$ 79",
    period: "/mês",
    description: "Para equipes pequenas",
    popular: true,
    features: {
      users: "5 usuários",
      storage: "100 GB",
      support: "Chat e Email",
      analytics: true,
      api: true,
      customDomain: true,
      ssl: true,
      backup: "Diário",
    },
  },
  {
    name: "Empresarial",
    price: "R$ 199",
    period: "/mês",
    description: "Para grandes equipes",
    popular: false,
    features: {
      users: "Ilimitado",
      storage: "1 TB",
      support: "Telefone 24/7",
      analytics: true,
      api: true,
      customDomain: true,
      ssl: true,
      backup: "Tempo real",
    },
  },
]

const featureLabels = {
  users: "Usuários",
  storage: "Armazenamento",
  support: "Suporte",
  analytics: "Analytics",
  api: "API Access",
  customDomain: "Domínio Personalizado",
  ssl: "Certificado SSL",
  backup: "Backup",
}

const getFeatureIcon = (key: string) => {
  const icons = {
    users: Users,
    storage: Shield,
    support: Zap,
    analytics: Star,
    api: Zap,
    customDomain: Shield,
    ssl: Shield,
    backup: Shield,
  }
  const Icon = icons[key as keyof typeof icons] || Shield
  return <Icon className="w-4 h-4" />
}

export default function Component() {
  return (
    <div className="container mx-auto p-6 space-y-8">
      <div className="text-center space-y-4">
        <h1 className="text-3xl font-bold">Compare Nossos Planos</h1>
        <p className="text-muted-foreground max-w-2xl mx-auto">
          Escolha o plano ideal para suas necessidades. Todos os planos incluem recursos essenciais para o seu sucesso.
        </p>
      </div>

      {/* Cards de Planos - Mobile First */}
      <div className="grid gap-6 md:grid-cols-3 lg:hidden">
        {plans.map((plan) => (
          <Card key={plan.name} className={`relative ${plan.popular ? "border-primary shadow-lg" : ""}`}>
            {plan.popular && (
              <Badge className="absolute -top-2 left-1/2 transform -translate-x-1/2">Mais Popular</Badge>
            )}
            <CardHeader className="text-center">
              <CardTitle className="text-xl">{plan.name}</CardTitle>
              <CardDescription>{plan.description}</CardDescription>
              <div className="flex items-baseline justify-center gap-1">
                <span className="text-3xl font-bold">{plan.price}</span>
                <span className="text-muted-foreground">{plan.period}</span>
              </div>
            </CardHeader>
            <CardContent className="space-y-4">
              {Object.entries(plan.features).map(([key, value]) => (
                <div key={key} className="flex items-center gap-3">
                  {getFeatureIcon(key)}
                  <span className="text-sm font-medium">{featureLabels[key as keyof typeof featureLabels]}</span>
                  <div className="ml-auto">
                    {typeof value === "boolean" ? (
                      value ? (
                        <Check className="w-4 h-4 text-green-500" />
                      ) : (
                        <X className="w-4 h-4 text-red-500" />
                      )
                    ) : (
                      <span className="text-sm text-muted-foreground">{value}</span>
                    )}
                  </div>
                </div>
              ))}
            </CardContent>
            <CardFooter>
              <Button className="w-full" variant={plan.popular ? "default" : "outline"}>
                Escolher Plano
              </Button>
            </CardFooter>
          </Card>
        ))}
      </div>

      {/* Tabela de Comparação - Desktop */}
      <div className="hidden lg:block">
        <div className="border rounded-lg overflow-hidden">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead className="w-48">Recursos</TableHead>
                {plans.map((plan) => (
                  <TableHead key={plan.name} className="text-center relative">
                    {plan.popular && (
                      <Badge className="absolute -top-2 left-1/2 transform -translate-x-1/2 z-10">Mais Popular</Badge>
                    )}
                    <div className="space-y-2 py-4">
                      <div className="font-semibold text-lg">{plan.name}</div>
                      <div className="text-sm text-muted-foreground">{plan.description}</div>
                      <div className="flex items-baseline justify-center gap-1">
                        <span className="text-2xl font-bold">{plan.price}</span>
                        <span className="text-sm text-muted-foreground">{plan.period}</span>
                      </div>
                      <Button className="w-full mt-2" variant={plan.popular ? "default" : "outline"}>
                        Escolher Plano
                      </Button>
                    </div>
                  </TableHead>
                ))}
              </TableRow>
            </TableHeader>
            <TableBody>
              {Object.entries(featureLabels).map(([key, label]) => (
                <TableRow key={key}>
                  <TableCell className="font-medium">
                    <div className="flex items-center gap-2">
                      {getFeatureIcon(key)}
                      {label}
                    </div>
                  </TableCell>
                  {plans.map((plan) => (
                    <TableCell key={`${plan.name}-${key}`} className="text-center">
                      {typeof plan.features[key as keyof typeof plan.features] === "boolean" ? (
                        plan.features[key as keyof typeof plan.features] ? (
                          <Check className="w-5 h-5 text-green-500 mx-auto" />
                        ) : (
                          <X className="w-5 h-5 text-red-500 mx-auto" />
                        )
                      ) : (
                        <span className="font-medium">
                          {plan.features[key as keyof typeof plan.features] as string}
                        </span>
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      </div>

      {/* Seção de FAQ ou informações adicionais */}
      <div className="text-center space-y-4 pt-8">
        <h2 className="text-xl font-semibold">Precisa de mais informações?</h2>
        <p className="text-muted-foreground">
          Entre em contato conosco para planos personalizados ou tire suas dúvidas.
        </p>
        <div className="flex gap-4 justify-center">
          <Button variant="outline">Falar com Vendas</Button>
          <Button variant="ghost">Ver FAQ</Button>
        </div>
      </div>
    </div>
  )
}
