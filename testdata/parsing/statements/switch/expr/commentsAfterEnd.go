package main
func main() {
    switch e {
        case a:
            fallthrough
        default:
            return 1
    }
    //adf
    //adf
    //adf
}
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  FunctionDeclaration(main)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('main')
    PsiElement(()('(')
    PsiElement())(')')
    PsiWhiteSpace(' ')
    BlockStmtImpl
      PsiElement({)('{')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      SwitchExprStmt
        PsiElement(KEYWORD_SWITCH)('switch')
        PsiWhiteSpace(' ')
        LiteralExpressionImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('e')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('        ')
        SwitchExprCase
          PsiElement(KEYWORD_CASE)('case')
          PsiWhiteSpace(' ')
          LiteralExpressionImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('a')
          PsiElement(:)(':')
          PsiWhiteSpace('\n')
          PsiWhiteSpace('            ')
          FallthroughStmt
            PsiElement(KEYWORD_FALLTHROUGH)('fallthrough')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('        ')
        SwitchExprCase
          PsiElement(KEYWORD_DEFAULT)('default')
          PsiElement(:)(':')
          PsiWhiteSpace('\n')
          PsiWhiteSpace('            ')
          ReturnStmtImpl
            PsiElement(KEYWORD_RETURN)('return')
            PsiWhiteSpace(' ')
            LiteralExpressionImpl
              LiteralIntegerImpl
                PsiElement(LITERAL_INT)('1')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('    ')
        PsiElement(})('}')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      PsiComment(SL_COMMENT)('//adf')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      PsiComment(SL_COMMENT)('//adf')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      PsiComment(SL_COMMENT)('//adf')
      PsiWhiteSpace('\n')
      PsiElement(})('}')
